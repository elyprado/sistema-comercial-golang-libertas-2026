const API = '/fornecedores';
let editandoId = null;
const modalFornecedor = new bootstrap.Modal(document.getElementById('modalFornecedor'));

const campos = ['txtNome', 'txtCnpj', 'txtTelefone', 'txtEmail',
  'txtLogradouro', 'txtNumero', 'txtBairro', 'txtCidade', 'txtUf'];

function getToken() {
  return sessionStorage.getItem('token') || localStorage.getItem('token') || '';
}

function authHeaders(extra = {}) {
  return { 'Authorization': 'Bearer ' + getToken(), ...extra };
}

function limparForm() {
  campos.forEach(id => {
    document.getElementById(id).value = '';
    limparErro(id);
  });
}

function preencherForm(f) {
  document.getElementById('txtNome').value = f.nome ?? '';
  document.getElementById('txtCnpj').value = formatarCnpj(f.cnpj ?? '');
  document.getElementById('txtTelefone').value = formatarTelefone(f.telefone ?? '');
  document.getElementById('txtEmail').value = f.email ?? '';
  document.getElementById('txtLogradouro').value = f.logradouro ?? '';
  document.getElementById('txtNumero').value = f.numero ?? '';
  document.getElementById('txtBairro').value = f.bairro ?? '';
  document.getElementById('txtCidade').value = f.cidade ?? '';
  document.getElementById('txtUf').value = f.uf ?? '';
  campos.forEach(limparErro);
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
    const resp = await fetch(`${API}/${id}`, {
      headers: authHeaders()
    });
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
    const resp = await fetch(`${API}/${id}`, {
      method: 'DELETE',
      headers: authHeaders()
    });
    if (!resp.ok) throw new Error(await resp.text());
    alert('Fornecedor excluído com sucesso!');
    carregaDados();
  } catch (e) {
    alert('Erro ao excluir fornecedor: ' + e.message);
  }
}

function marcarErro(id) {
  document.getElementById(id).classList.add('is-invalid');
}

function limparErro(id) {
  document.getElementById(id).classList.remove('is-invalid');
}

function cnpjValido(cnpj) {
  return /^\d{14}$/.test(cnpj);
}

function emailValido(email) {
  if (!email) return true;
  return /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email);
}

function validarForm(dados) {
  let valido = true;

  if (!dados.nome) {
    marcarErro('txtNome');
    valido = false;
  }

  if (!cnpjValido(dados.cnpj)) {
    marcarErro('txtCnpj');
    valido = false;
  }

  if (dados.telefone && ![10, 11].includes(dados.telefone.length)) {
    marcarErro('txtTelefone');
    valido = false;
  }

  if (!emailValido(dados.email)) {
    marcarErro('txtEmail');
    valido = false;
  }

  return valido;
}

function getForm() {
  return {
    nome: document.getElementById('txtNome').value.trim(),
    cnpj: document.getElementById('txtCnpj').value.replace(/\D/g, ''),
    telefone: document.getElementById('txtTelefone').value.replace(/\D/g, ''),
    email: document.getElementById('txtEmail').value.trim(),
    logradouro: document.getElementById('txtLogradouro').value.trim(),
    numero: document.getElementById('txtNumero').value.replace(/\D/g, ''),
    bairro: document.getElementById('txtBairro').value.trim(),
    cidade: document.getElementById('txtCidade').value.trim(),
    uf: document.getElementById('txtUf').value.trim()
  };
}

async function salvar() {
  campos.forEach(limparErro);
  const fornecedor = getForm();

  if (!validarForm(fornecedor)) return;

  const metodo = editandoId ? 'PUT' : 'POST';
  const url = editandoId ? `${API}/${editandoId}` : API;

  try {
    const resp = await fetch(url, {
      method: metodo,
      headers: authHeaders({ 'Content-Type': 'application/json' }),
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

function formatarCnpj(cnpj) {
  const v = (cnpj || '').replace(/\D/g, '');
  if (v.length !== 14) return cnpj || '';
  return v.replace(/(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})/, '$1.$2.$3/$4-$5');
}

function formatarTelefone(tel) {
  const v = (tel || '').replace(/\D/g, '');
  if (v.length === 11) return v.replace(/(\d{2})(\d{5})(\d{4})/, '($1) $2-$3');
  if (v.length === 10) return v.replace(/(\d{2})(\d{4})(\d{4})/, '($1) $2-$3');
  return tel || '';
}

async function carregaDados() {
  const tbody = document.getElementById('dados');
  tbody.innerHTML = '<tr><td colspan="7" class="text-center text-muted py-4" style="max-width:none;">Carregando...</td></tr>';

  try {
    const resp = await fetch(API, {
      headers: authHeaders()
    });
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
      const cnpjFmt = formatarCnpj(f.cnpj);
      const telFmt = formatarTelefone(f.telefone);
      tbody.innerHTML += `
        <tr>
          <td class="ps-3 text-muted">${f.idfornecedor ?? '-'}</td>
          <td class="fw-semibold" title="${f.nome ?? ''}">${f.nome ?? '-'}</td>
          <td title="${cnpjFmt}">${cnpjFmt || '-'}</td>
          <td title="${telFmt}">${telFmt || '-'}</td>
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
    const isToken = e.message.toLowerCase().includes('token') || e.message.toLowerCase().includes('unauthorized') || e.message === 'Failed to fetch' || e.message.includes('NetworkError');
    tbody.innerHTML = `
      <tr>
        <td colspan="7" style="max-width:none; padding: 48px 24px;">
          <div class="d-flex flex-column align-items-center text-center">
            <div style="width:56px;height:56px;background:#fff1f1;border-radius:16px;display:flex;align-items:center;justify-content:center;margin-bottom:16px;">
              <svg width="28" height="28" fill="none" viewBox="0 0 24 24" stroke="#dc3545" stroke-width="2">
                <circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/>
              </svg>
            </div>
            <h6 class="fw-semibold text-dark mb-1">${isToken ? 'Falha de Conexão' : 'Erro ao carregar dados'}</h6>
            <p class="text-muted mb-0" style="font-size:0.83rem;max-width:320px;">
              ${isToken ? 'Token não encontrado. Faça o login.' : e.message}
            </p>
          </div>
        </td>
      </tr>`;
    document.getElementById('rodape').style.display = 'none';
  }
}

document.getElementById('txtCnpj').addEventListener('input', function () {
  let v = this.value.replace(/\D/g, '').slice(0, 14);
  v = v.replace(/(\d{2})(\d)/, '$1.$2')
    .replace(/(\d{3})(\d)/, '$1.$2')
    .replace(/(\d{3})(\d)/, '$1/$2')
    .replace(/(\d{4})(\d{1,2})$/, '$1-$2');
  this.value = v;
  limparErro('txtCnpj');
});

document.getElementById('txtTelefone').addEventListener('input', function () {
  let v = this.value.replace(/\D/g, '').slice(0, 11);
  if (v.length > 10) {
    v = v.replace(/(\d{2})(\d{5})(\d{0,4})/, '($1) $2-$3');
  } else if (v.length > 6) {
    v = v.replace(/(\d{2})(\d{4})(\d{0,4})/, '($1) $2-$3');
  } else if (v.length > 2) {
    v = v.replace(/(\d{2})(\d{0,4})/, '($1) $2');
  } else if (v.length > 0) {
    v = v.replace(/(\d{0,2})/, '($1');
  }
  this.value = v;
  limparErro('txtTelefone');
});

document.getElementById('txtNumero').addEventListener('input', function () {
  this.value = this.value.replace(/\D/g, '').slice(0, 10);
});

document.getElementById('txtNome').addEventListener('input', () => limparErro('txtNome'));
document.getElementById('txtEmail').addEventListener('input', () => limparErro('txtEmail'));

carregaDados();