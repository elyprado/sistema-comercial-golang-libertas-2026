const API = 'http://localhost:8080/clientes';
let editandoId = null;

// GERENCIAMENTO DAS TELAS
function abrirFormulario(cliente = null) {
    // Esconde a lista e mostra o form
    document.getElementById('view-list').classList.add('hidden');
    document.getElementById('view-form').classList.remove('hidden');

    if (cliente) {
        // Modo edição
        editandoId = cliente.idcliente;
        document.getElementById('breadcrumb-current').textContent = 'Editar Cadastro';

        // Preenche os campos
        ['nome', 'email', 'cpf', 'telefone', 'logradouro', 'numero', 'bairro', 'cidade', 'uf']
            .forEach(f => {
                const el = document.getElementById('f-' + f);
                if (el) el.value = cliente[f] || '';
                clearError(f);
            });
    } else {
        // Cadastro novo
        editandoId = null;
        document.getElementById('breadcrumb-current').textContent = 'Novo Cadastro';
        limparFormulario();
    }
}

function voltarParaLista() {
    // Esconde o form e mostra a lista
    document.getElementById('view-form').classList.add('hidden');
    document.getElementById('view-list').classList.remove('hidden');
    carregarClientes(); // Atualiza a lista sempre que voltar
}

// UTILIDADES
function showMsg(txt, ok = true) {
    const m = document.getElementById('msg');
    m.textContent = txt;

    if (ok) {
        m.className = 'fixed top-6 right-6 px-5 py-3.5 rounded-lg text-sm font-medium border shadow-lg transition-all duration-300 z-50 bg-green-50 text-green-800 border-green-200';
    } else {
        m.className = 'fixed top-6 right-6 px-5 py-3.5 rounded-lg text-sm font-medium border shadow-lg transition-all duration-300 z-50 bg-red-50 text-red-800 border-red-200';
    }

    m.classList.remove('hidden');
    setTimeout(() => m.classList.add('hidden'), 4000);
}

// Visual de validação
function setError(id) {
    const input = document.getElementById(`f-${id}`);
    const msg = document.getElementById(`err-${id}`);
    input.classList.remove('border-gray-300', 'focus:ring-gray-900');
    input.classList.add('border-red-500', 'focus:ring-red-500', 'bg-red-50');
    if (msg) msg.classList.remove('hidden');
}

function clearError(id) {
    const input = document.getElementById(`f-${id}`);
    const msg = document.getElementById(`err-${id}`);
    input.classList.remove('border-red-500', 'focus:ring-red-500', 'bg-red-50');
    input.classList.add('border-gray-300', 'focus:ring-gray-900');
    if (msg) msg.classList.add('hidden');
}

document.getElementById('f-nome').addEventListener('input', () => clearError('nome'));
document.getElementById('f-cpf').addEventListener('input', () => clearError('cpf'));

function limparFormulario() {
    ['nome', 'email', 'cpf', 'telefone', 'logradouro', 'numero', 'bairro', 'cidade', 'uf']
        .forEach(f => {
            document.getElementById('f-' + f).value = '';
            clearError(f);
        });
}

function getForm() {
    return {
        nome: document.getElementById('f-nome').value.trim(),
        email: document.getElementById('f-email').value.trim(),
        cpf: document.getElementById('f-cpf').value.replace(/\D/g, ''),
        telefone: document.getElementById('f-telefone').value.trim(),
        logradouro: document.getElementById('f-logradouro').value.trim(),
        numero: document.getElementById('f-numero').value.trim(),
        bairro: document.getElementById('f-bairro').value.trim(),
        cidade: document.getElementById('f-cidade').value.trim(),
        uf: document.getElementById('f-uf').value.trim()
    };
}

// COMUNICAÇÃO API
async function salvar() {
    const dados = getForm();
    let isValido = true;

    if (!dados.nome) { setError('nome'); isValido = false; }
    if (!dados.cpf) { setError('cpf'); isValido = false; }
    if (!isValido) return;

    try {
        const method = editandoId ? 'PUT' : 'POST';
        const url = editandoId ? `${API}/${editandoId}` : API;
        const r = await fetch(url, {
            method,
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(dados)
        });

        if (!r.ok) throw new Error(await r.text());

        showMsg(editandoId ? 'Cadastro atualizado com sucesso!' : 'Cliente registrado com sucesso!');
        voltarParaLista(); // Se salvou com sucesso, volta pra tabela
    } catch (e) {
        showMsg('Erro na requisição: ' + e.message, false);
    }
}

async function carregarClientes() {
    const wrap = document.getElementById('tabela-wrap');

    wrap.innerHTML = `
      <div class="flex flex-col items-center justify-center py-24 animate-pulse">
        <div class="w-10 h-10 rounded-full border-4 border-gray-100 border-t-gray-900 animate-spin mb-4"></div>
        <p class="text-gray-500 font-medium text-sm">Buscando dados...</p>
      </div>`;

    try {
        const r = await fetch(API);
        if (!r.ok) throw new Error(await r.text());
        const lista = await r.json();

        // Tela vazia
        if (!lista || lista.length === 0) {
            wrap.innerHTML = `
              <div class="flex flex-col items-center justify-center py-24 text-center">
                <div class="w-16 h-16 bg-gray-50 rounded-full flex items-center justify-center text-gray-400 mb-4">
                  <svg class="w-8 h-8" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"></path></svg>
                </div>
                <h3 class="text-base font-medium text-gray-900">Nenhum registro</h3>
                <p class="text-gray-500 mt-1 text-sm">Sua base de clientes está vazia.</p>
              </div>`;
            return;
        }

        wrap.innerHTML = `
          <div class="overflow-x-auto">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="border-b border-gray-200 bg-gray-50/50">
                  <th class="px-6 py-4 text-sm font-semibold text-gray-700">ID</th>
                  <th class="px-6 py-4 text-sm font-semibold text-gray-700">Cliente</th>
                  <th class="px-6 py-4 text-sm font-semibold text-gray-700">Contato</th>
                  <th class="px-6 py-4 text-sm font-semibold text-gray-700">Localidade</th>
                  <th class="px-6 py-4 text-sm font-semibold text-gray-700 text-right">Ações</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-100 bg-white">
                ${lista.map(c => `
                  <tr class="hover:bg-gray-50/50 transition-colors group">
                    <td class="px-6 py-4 text-sm text-gray-500">${c.idcliente || '-'}</td>
                    <td class="px-6 py-4">
                      <div class="text-sm font-medium text-gray-900">${c.nome || ''}</div>
                      <div class="text-xs text-gray-500 mt-0.5">${c.cpf || 'Sem CPF'}</div>
                    </td>
                    <td class="px-6 py-4">
                      <div class="text-sm text-gray-700">${c.telefone || '-'}</div>
                      <div class="text-xs text-gray-500 mt-0.5">${c.email || '-'}</div>
                    </td>
                    <td class="px-6 py-4">
                      <div class="text-sm text-gray-700">${c.cidade || '-'}</div>
                      <div class="text-xs text-gray-500 mt-0.5">${c.uf || ''}</div>
                    </td>
                    <td class="px-6 py-4 text-right text-sm font-medium">
                      <div class="flex justify-end gap-3 opacity-0 group-hover:opacity-100 transition-opacity">
                        <button onclick='abrirFormulario(${JSON.stringify(c)})' class="text-gray-400 hover:text-black transition-colors">Editar</button>
                        <button onclick="deletar(${c.idcliente})" class="text-gray-400 hover:text-red-600 transition-colors">Excluir</button>
                      </div>
                    </td>
                  </tr>`).join('')}
              </tbody>
            </table>
          </div>`;
    } catch (e) {
        // Mensagem de erro
        wrap.innerHTML = `
          <div class="flex flex-col items-center justify-center py-20 text-center">
            <div class="w-16 h-16 bg-red-50 rounded-2xl flex items-center justify-center text-red-500 mb-4">
              <svg class="w-8 h-8" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
            </div>
            <h3 class="text-base font-medium text-gray-900">Falha de Conexão</h3>
            <p class="text-gray-500 mt-1 max-w-sm text-sm">${e.message}</p>
          </div>`;
    }
}

async function deletar(id) {
    if (!confirm('Excluir permanentemente este registro?')) return;
    try {
        const r = await fetch(`${API}/${id}`, { method: 'DELETE' });
        if (!r.ok) throw new Error(await r.text());
        showMsg('Registro excluído com sucesso.');
        carregarClientes();
    } catch (e) {
        showMsg('Erro ao excluir: ' + e.message, false);
    }
}

// Máscaras de input
document.getElementById('f-cpf').addEventListener('input', function () {
    let v = this.value.replace(/\D/g, '').slice(0, 11);
    v = v.replace(/(\d{3})(\d)/, '$1.$2')
        .replace(/(\d{3})(\d)/, '$1.$2')
        .replace(/(\d{3})(\d{1,2})$/, '$1-$2');
    this.value = v;
});

document.getElementById('f-uf').addEventListener('input', function () {
    this.value = this.value.toUpperCase().replace(/[^A-Z]/g, '');
});

// Carrega a lista assim que abrir a página
carregarClientes();