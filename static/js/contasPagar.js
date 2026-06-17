// Instancia o modal do Bootstrap seguindo o padrão do usuarios.html
const modalElement = document.getElementById('modalConta');
const modalConta = window.bootstrap ? new window.bootstrap.Modal(modalElement) : new bootstrap.Modal(modalElement);

document.addEventListener("DOMContentLoaded", () => {
    carregaDados();

    // Listener para aplicar o filtro por vencimento imediatamente ao alterar a data
    document.getElementById("filtroVencimento").addEventListener("change", carregaDados);
});

// 1. REQUISITO: Listar registros e filtrar por vencimento (GET)
async function carregaDados() {
    const tbody = document.getElementById("dadosContas");
    tbody.innerHTML = `<tr><td colspan="7" class="text-center py-3 help-text">Carregando dados...</td></tr>`;
    
    const filtroData = document.getElementById("filtroVencimento").value;
    let url = "/contas-pagar";

    // Se houver filtro preenchido na tela, passa via Query Parameter para o Go
    if (filtroData) {
        url += `?vencimento=${filtroData}`;
    }

    try {
        const resp = await fetch(url);
        if (!resp.ok) throw new Error();
        const data = await resp.json();
        
        tbody.innerHTML = "";

        if (!data || data.length === 0) {
            tbody.innerHTML = `<tr><td colspan="7" class="text-center py-4 help-text">Nenhuma conta encontrada.</td></tr>`;
            return;
        }

        for (const c of data) {
            // REQUISITO: Exibir em cor diferente os itens que já foram pagos
            // Verifica se a string existe e se não é a "data zero" padrão do Go (0001-01-01)
            const estaPaga = c.data_pagamento && !c.data_pagamento.startsWith("0001-01-01");
            const classeCor = estaPaga ? "conta-paga" : "conta-pendente";

            // Formatação de datas para exibição legível em formato BR
            const dtEmissao = formatarDataBR(c.data);
            const dtVenc = formatarDataBR(c.data_vencimento);
            const dtPagto = estaPaga ? formatarDataBR(c.data_pagamento) : `<span class="badge bg-warning text-dark">Em aberto</span>`;

            // REQUISITO: Botão de baixar item (Apenas para contas em aberto)
            const acaoBotao = !estaPaga 
                ? `<button class="btn btn-sm btn-success fw-bold px-3" onclick="baixar(${c.id_contas_pagar})">💳 Baixar</button>` 
                : `<span class="text-success fw-bold small">✓ Pago</span>`;

            tbody.innerHTML += `
                <tr class="${classeCor}">
                    <td><strong>#${c.id_contas_pagar}</strong></td>
                    <td>Fornecedor ${c.id_forncedor}</td> <td>${dtEmissao}</td>
                    <td>${dtVenc}</td>
                    <td><strong>R$ ${parseFloat(c.valor).toFixed(2)}</strong></td>
                    <td>${dtPagto}</td>
                    <td class="text-center">${acaoBotao}</td>
                </tr>
            `;
        }
    } catch (error) {
        tbody.innerHTML = `<tr><td colspan="7" class="text-center py-4 text-danger fw-bold">Erro ao carregar dados do servidor Go.</td></tr>`;
    }
}

// 2. REQUISITO: Inserir nova conta a pagar (POST)
async function salvar() {
    const emissao = document.getElementById("txtDataEmissao").value;
    const vencimento = document.getElementById("txtVencimento").value;
    const fornecedor = document.getElementById("txtFornecedor").value;
    const valor = document.getElementById("txtValor").value;

    if (!fornecedor || !emissao || !vencimento || !valor) {
        alert("Por favor, preencha todos os campos do formulário.");
        return;
    }

    // Monta o JSON casado perfeitamente com o ContasPagarRequest do Go
    // Nota: O sufixo "T00:00:00Z" é obrigatório para que o tipo time.Time do Go consiga ler a string de data sem dar erro.
    let contaPayload = {
        id_forncedor: parseInt(fornecedor), // Mantido erro de digitação do struct do grupo
        data: `${emissao}T00:00:00Z`, 
        data_vencimento: `${vencimento}T00:00:00Z`,
        valor: parseFloat(valor)
    };

    try {
        const resp = await fetch("/contas-pagar", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(contaPayload)
        });

        if (resp.ok) {
            alert("Conta lançada com sucesso!");
            modalConta.hide();
            carregaDados();
        } else {
            alert("Erro ao salvar o lançamento no backend.");
        }
    } catch (error) {
        alert("Falha de comunicação com o servidor.");
    }
}

// 3. REQUISITO: Baixar item salvando o atributo data_pagamento (PUT)
async function baixar(idContasPagar) {
    if (!confirm("Deseja realmente confirmar o pagamento e baixar este título?")) return;

    // Pega a data atual do computador do usuário no formato YYYY-MM-DD
    const hoje = new Date().toISOString().split('T')[0];
    
    // Payload enviado para o método de baixa do seu controller
    const payloadBaixa = { data_pagamento: `${hoje}T00:00:00Z` };

    try {
        const resp = await fetch(`/contas-pagar/${idContasPagar}/baixar`, {
            method: "PUT",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(payloadBaixa)
        });

        if (resp.ok) {
            alert("Baixa processada com sucesso!");
            carregaDados();
        } else {
            alert("Não foi possível processar a baixa.");
        }
    } catch (error) {
        alert("Erro de rede ao processar a baixa.");
    }
}

// Funções auxiliares de controle da interface
function abrirModalNovo() {
    document.getElementById("txtFornecedor").value = "";
    document.getElementById("txtDataEmissao").value = "";
    document.getElementById("txtVencimento").value = "";
    document.getElementById("txtValor").value = "";
    modalConta.show();
}

function limparFiltro() {
    document.getElementById("filtroVencimento").value = "";
    carregaDados();
}

function formatarDataBR(dataSql) {
    if (!dataSql || dataSql.startsWith("0001-01-01")) return "";
    const dataApenas = dataSql.split("T")[0]; 
    const partes = dataApenas.split("-");
    if (partes.length !== 3) return dataSql;
    return `${partes[2]}/${partes[1]}/${partes[0]}`;
}