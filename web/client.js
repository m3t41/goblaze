// Copyright 2026 Daniel\n// Licensed under the GNU Affero General Public License v3.0.\n// Copying or distributing this file requires compliance with AGPLv3.\n
const root = document.getElementById("root");

function createNode(n) {
  // Textnode
  if (n.type === 1) {
    return document.createTextNode(n.text || "");
  }

  // Elementnode
  const el = document.createElement(n.tag || "div");

  if (n.props) {
    for (const [k, v] of Object.entries(n.props)) {
      el.setAttribute(k, v);
    }
  }

  if (n.children) {
    for (const c of n.children) {
      el.appendChild(createNode(c));
    }
  }

  return el;
}

function getNodeByPath(path) {
  let node = root;
  for (const idx of path) {
    node = node.childNodes[idx];
  }
  return node;
}

function applyPatch(p) {
  // 🔥 Root-Replace (path = [])
  if (!p.path || p.path.length === 0) {
    root.innerHTML = "";
    root.appendChild(createNode(p.node));
    return;
  }

  const node = getNodeByPath(p.path);

  switch (p.op) {
    case "setText":
      node.textContent = p.value;
      break;

    case "setAttr":
      node.setAttribute(p.name, p.value);
      break;

    case "removeAttr":
      node.removeAttribute(p.name);
      break;

    case "insert": {
      const parent = getNodeByPath(p.path.slice(0, -1));
      parent.appendChild(createNode(p.node));
      break;
    }

    case "remove":
      node.remove();
      break;

    case "replace":
      node.replaceWith(createNode(p.node));
      break;
  }
}

// WebSocket verbinden
const ws = new WebSocket(`ws://${location.host}/_goblaze`);

ws.onmessage = ev => {
  const patches = JSON.parse(ev.data);
  if (!Array.isArray(patches)) return;
  patches.forEach(applyPatch);
};

// Events an Server senden
document.addEventListener("click", e => {
  if (e.target.tagName === "BUTTON") {
    const msg = e.target.getAttribute("onclick");
    if (msg) ws.send(msg);
  }
});
