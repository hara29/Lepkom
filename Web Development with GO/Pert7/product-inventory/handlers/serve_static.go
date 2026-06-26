package handlers

import "net/http"

func RedirectHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func ServeLogin(w http.ResponseWriter, r *http.Request) {
	writeHTML(w, loginHTML)
}

func ServeProductsPage(w http.ResponseWriter, r *http.Request) {
	writeHTML(w, productsHTML)
}

func ServeUsersPage(w http.ResponseWriter, r *http.Request) {
	writeHTML(w, usersHTML)
}

func writeHTML(w http.ResponseWriter, html string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(html))
}

const baseStyle = `<style>
  body { font-family: Arial, sans-serif; margin: 0; background: #f6f7f9; color: #17202a; }
  header { background: #0f766e; color: white; padding: 18px 28px; display: flex; justify-content: space-between; align-items: center; gap: 16px; }
  nav a { color: white; margin-left: 14px; text-decoration: none; font-weight: 700; }
  main { max-width: 960px; margin: 24px auto; padding: 0 16px; }
  section { background: white; border: 1px solid #dde3ea; border-radius: 8px; padding: 18px; margin-bottom: 16px; }
  input, select, button { padding: 10px; margin: 4px 0; border: 1px solid #cbd5df; border-radius: 6px; }
  button { background: #0f766e; color: white; cursor: pointer; border: 0; }
  table { width: 100%; border-collapse: collapse; margin-top: 12px; }
  th, td { border-bottom: 1px solid #e5e9ef; padding: 10px; text-align: left; }
  .table-wrap { width: 100%; overflow-x: auto; }
  .users-table { table-layout: fixed; }
  .users-table th:nth-child(1), .users-table td:nth-child(1) { width: 52px; }
  .users-table th:nth-child(2), .users-table td:nth-child(2) { width: 140px; }
  .users-table th:nth-child(3), .users-table td:nth-child(3) { width: 88px; }
  .token-cell { display: grid; grid-template-columns: minmax(0, 1fr) auto; gap: 8px; align-items: start; }
  .token-text { display: block; max-height: 76px; overflow: auto; overflow-wrap: anywhere; word-break: break-word; background: #f1f5f9; border: 1px solid #d8e0ea; border-radius: 6px; padding: 8px; font-family: ui-monospace, SFMono-Regular, Menlo, Consolas, monospace; font-size: 12px; line-height: 1.45; }
  .copy-btn { white-space: nowrap; padding: 8px 10px; }
  .row { display: grid; grid-template-columns: repeat(4, minmax(0, 1fr)); gap: 8px; }
  .actions { display: flex; gap: 8px; }
  pre { background: #101828; color: #d1fadf; padding: 12px; overflow: auto; border-radius: 6px; }
  @media (max-width: 720px) { header { align-items: flex-start; flex-direction: column; } nav a { margin-left: 0; margin-right: 12px; } .row { grid-template-columns: 1fr; } }
</style>`

const loginHTML = `<!doctype html>
<html lang="id">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Login - Product Inventory</title>` + baseStyle + `
</head>
<body>
  <header>
    <h1>Login</h1>
    <nav><a href="/products">Produk</a><a href="/users">User Management</a></nav>
  </header>
  <main>
    <section>
      <div class="row">
        <input id="username" placeholder="Username" value="admin">
        <input id="password" placeholder="Password" type="password" value="password123">
        <button onclick="login()">Login</button>
      </div>
      <pre id="output"></pre>
    </section>
  </main>
  <script>
    async function login() {
      const res = await fetch('/api/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ username: username.value, password: password.value })
      });
      const json = await res.json();
      output.textContent = JSON.stringify(json, null, 2);
      if (json.data && json.data.token) {
        localStorage.setItem('token', json.data.token);
        location.href = '/products';
      }
    }
  </script>
</body>
</html>`

const productsHTML = `<!doctype html>
<html lang="id">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Produk - Product Inventory</title>` + baseStyle + `
</head>
<body onload="loadProducts()">
  <header>
    <h1>Produk</h1>
    <nav><a href="/login">Login</a><a href="/users">User Management</a></nav>
  </header>
  <main>
    <section>
      <div class="row">
        <input id="productName" placeholder="Nama produk">
        <input id="productStock" placeholder="Stok" type="number">
        <input id="productPrice" placeholder="Harga" type="number">
        <button onclick="saveProduct()">Simpan</button>
      </div>
      <table>
        <thead><tr><th>ID</th><th>Nama</th><th>Stok</th><th>Harga</th><th>Aksi</th></tr></thead>
        <tbody id="products"></tbody>
      </table>
      <pre id="output"></pre>
    </section>
  </main>
  <script>
    let token = localStorage.getItem('token') || '';
    let editingID = null;

    function authHeaders() {
      return { 'Content-Type': 'application/json', 'Authorization': 'Bearer ' + token };
    }

    async function loadProducts() {
      const res = await fetch('/api/products', { headers: authHeaders() });
      const json = await res.json();
      output.textContent = JSON.stringify(json, null, 2);
      products.innerHTML = '';
      (json.data || []).forEach((p, index) => {
        const row = document.createElement('tr');
        row.innerHTML = '<td>' + p.id + '</td><td>' + p.name + '</td><td>' + p.stock + '</td><td>' + p.price + '</td><td class="actions"><button type="button" data-index="' + index + '" class="edit-btn">Edit</button><button type="button" data-id="' + p.id + '" class="delete-btn">Hapus</button></td>';
        products.appendChild(row);
      });

      document.querySelectorAll('.edit-btn').forEach(button => {
        button.addEventListener('click', () => {
          const product = json.data[Number(button.dataset.index)];
          editProduct(product);
        });
      });

      document.querySelectorAll('.delete-btn').forEach(button => {
        button.addEventListener('click', () => deleteProduct(Number(button.dataset.id)));
      });
    }

    function editProduct(product) {
      editingID = product.id;
      productName.value = product.name;
      productStock.value = product.stock;
      productPrice.value = product.price;
    }

    async function saveProduct() {
      const body = JSON.stringify({ name: productName.value, stock: Number(productStock.value), price: Number(productPrice.value) });
      const url = editingID ? '/api/products/' + editingID : '/api/products';
      const method = editingID ? 'PUT' : 'POST';
      const res = await fetch(url, { method, headers: authHeaders(), body });
      output.textContent = JSON.stringify(await res.json(), null, 2);
      editingID = null;
      productName.value = '';
      productStock.value = '';
      productPrice.value = '';
      loadProducts();
    }

    async function deleteProduct(id) {
      const res = await fetch('/api/products/' + id, { method: 'DELETE', headers: authHeaders() });
      output.textContent = JSON.stringify(await res.json(), null, 2);
      loadProducts();
    }
  </script>
</body>
</html>`

const usersHTML = `<!doctype html>
<html lang="id">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>User Management - Product Inventory</title>` + baseStyle + `
</head>
<body onload="loadUsers()">
  <header>
    <h1>User Management</h1>
    <nav><a href="/login">Login</a><a href="/products">Produk</a></nav>
  </header>
  <main>
    <section>
      <div class="row">
        <input id="newUsername" placeholder="Username baru">
        <input id="newPassword" placeholder="Password baru" type="password">
        <select id="newRole"><option value="user">user</option><option value="admin">admin</option></select>
        <button onclick="createUser()">Tambah User</button>
      </div>
    </section>
    <section>
      <div class="row">
        <input id="tokenUser" placeholder="Username token" value="admin">
        <button onclick="generateToken()">Generate Token</button>
        <button onclick="loadUsers()">Refresh User</button>
      </div>
      <div class="table-wrap">
        <table class="users-table">
          <thead><tr><th>ID</th><th>Username</th><th>Role</th><th>API Token</th></tr></thead>
          <tbody id="users"></tbody>
        </table>
      </div>
      <pre id="output"></pre>
    </section>
  </main>
  <script>
    const token = localStorage.getItem('token') || '';

    function authHeaders() {
      return { 'Content-Type': 'application/json', 'Authorization': 'Bearer ' + token };
    }

    async function loadUsers() {
      const res = await fetch('/api/users', { headers: authHeaders() });
      const json = await res.json();
      output.textContent = JSON.stringify(json, null, 2);
      users.innerHTML = '';
      (json.data || []).forEach((u, index) => {
        const row = document.createElement('tr');
        row.innerHTML = '<td>' + u.id + '</td><td>' + u.username + '</td><td>' + u.role + '</td><td><div class="token-cell"><code class="token-text">' + (u.api_token || '-') + '</code>' + (u.api_token ? '<button type="button" class="copy-btn" data-index="' + index + '">Copy</button>' : '') + '</div></td>';
        users.appendChild(row);
      });

      document.querySelectorAll('.copy-btn').forEach(button => {
        button.addEventListener('click', async () => {
          const user = json.data[Number(button.dataset.index)];
          await navigator.clipboard.writeText(user.api_token || '');
          button.textContent = 'Copied';
          setTimeout(() => button.textContent = 'Copy', 1200);
        });
      });
    }

    async function createUser() {
      const res = await fetch('/api/register', {
        method: 'POST',
        headers: authHeaders(),
        body: JSON.stringify({ username: newUsername.value, password: newPassword.value, role: newRole.value })
      });
      output.textContent = JSON.stringify(await res.json(), null, 2);
      newUsername.value = '';
      newPassword.value = '';
      loadUsers();
    }

    async function generateToken() {
      const res = await fetch('/api/generate-token', {
        method: 'POST',
        headers: authHeaders(),
        body: JSON.stringify({ username: tokenUser.value })
      });
      output.textContent = JSON.stringify(await res.json(), null, 2);
      loadUsers();
    }
  </script>
</body>
</html>`
