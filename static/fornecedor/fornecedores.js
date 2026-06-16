const API = '/fornecedores';
let editandoId = null;
const modalFornecedor = new bootstrap.Modal(document.getElementById('modalFornecedor'));

function limparForm() {
  ['txtNome', 'txtCnpj', 'txtTelefone', 'txtEmail',
    'txtLogradouro', 'txtNumero', 'txtBairro', 'txtCidade', 'txtUf']
    .forEach(id => document.getElementById(id).value = '');
}

function preencherForm(f) {
  document.getElementById('txtNome').value = f.nome ?? '';
  document.getElementById('txtCnpj').value = f.cnpj ?? '';
  document.getElementById('txtTelefone').value = f.telefone ?? '';
  document.getElementById('txtEmail').value = f.email ?? '';
  document.getElementById('txtLogradouro').value = f.logradouro ?? '';
  document.getElementById('txtNumero').value = f.numero ?? '';
  document.getElementById('txtBairro').value = f.bairro ?? '';
  document.getElementById('txtCidade').value = f.cidade ?? '';
  document.getElementById('txtUf').value = f.uf ?? '';
}

function novo() {
  editandoId = null;
  limparForm();
  document.getElementById('modalFornecedorLabel').textContent = 'Novo Fornecedor';
  modalFornecedor.show();
}

async function editar(id) {
  try {
    editandoId = id;
    const resp = await fetch(`${API}/${id}`);
    if (!resp.ok) throw new Error(await resp.text());
    const f = await resp.json();
    preencherForm(f);
    document.getElementById('modalFornecedorLabel').textContent = 'Editar Fornecedor';
    modalFornecedor.show();
  } catch (e) {
    alert('Erro ao carregar fornecedor: ' + e.message);
  }
}

async function excluir(id) {
  if (!confirm('Deseja realmente excluir este fornecedor?')) return;
  try {
    const resp = await fetch(`${API}/${id}`, { method: 'DELETE' });
    if (!resp.ok) throw new Error(await resp.text());
    alert('Fornecedor excluído com sucesso!');
    carregaDados();
  } catch (e) {
    alert('Erro ao excluir fornecedor: ' + e.message);
  }
}

async function salvar() {
  const fornecedor = {
    nome: document.getElementById('txtNome').value,
    cnpj: document.getElementById('txtCnpj').value,
    telefone: document.getElementById('txtTelefone').value,
    email: document.getElementById('txtEmail').value,
    logradouro: document.getElementById('txtLogradouro').value,
    numero: document.getElementById('txtNumero').value,
    bairro: document.getElementById('txtBairro').value,
    cidade: document.getElementById('txtCidade').value,
    uf: document.getElementById('txtUf').value
  };

  const metodo = editandoId ? 'PUT' : 'POST';
  const url = editandoId ? `${API}/${editandoId}` : API;

  try {
    const resp = await fetch(url, {
      method: metodo,
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(fornecedor)
    });

    if (!resp.ok) throw new Error(await resp.text());

    alert('Fornecedor salvo com sucesso!');
    modalFornecedor.hide();
    carregaDados();
  } catch (e) {
    alert('Erro ao salvar fornecedor: ' + e.message);
  }
}

async function carregaDados() {
  const tbody = document.getElementById('dados');
  tbody.innerHTML = '<tr><td colspan="7" class="text-center text-muted py-4" style="max-width:none;">Carregando...</td></tr>';

  try {
    const resp = await fetch(API);
    if (!resp.ok) throw new Error(await resp.text());
    const data = await resp.json();

    if (!data || data.length === 0) {
      tbody.innerHTML = '<tr><td colspan="7" class="text-center text-muted py-4" style="max-width:none;">Nenhum fornecedor cadastrado.</td></tr>';
      document.getElementById('rodape').style.display = 'none';
      return;
    }

    tbody.innerHTML = '';
    for (const f of data) {
      const cidadeUf = [f.cidade, f.uf].filter(Boolean).join(' / ') || '-';
      tbody.innerHTML += `
        <tr>
          <td class="ps-3 text-muted">${f.idfornecedor ?? '-'}</td>
          <td class="fw-semibold" title="${f.nome ?? ''}">${f.nome ?? '-'}</td>
          <td title="${f.cnpj ?? ''}">${f.cnpj ?? '-'}</td>
          <td title="${f.telefone ?? ''}">${f.telefone ?? '-'}</td>
          <td class="text-muted" title="${f.email ?? ''}">${f.email ?? '-'}</td>
          <td title="${cidadeUf}">${cidadeUf}</td>
          <td class="text-center col-acoes">
            <button class="btn btn-sm btn-outline-secondary me-1" onclick="editar(${f.idfornecedor})">Editar</button>
            <button class="btn btn-sm btn-outline-danger" onclick="excluir(${f.idfornecedor})">Excluir</button>
          </td>
        </tr>`;
    }
    document.getElementById('totalRegistros').textContent = data.length;
    document.getElementById('rodape').style.display = 'flex';
  } catch (e) {
    tbody.innerHTML = `<tr><td colspan="7" class="text-center text-danger py-4" style="max-width:none;">Erro ao carregar dados: ${e.message}</td></tr>`;
    document.getElementById('rodape').style.display = 'none';
  }
}

carregaDados();
