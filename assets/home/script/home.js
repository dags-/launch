const view = {data: {modpacks: []}, launch: function(){}, selectPack: function(){}};

function onload() {
    Render.append("social", socials, SocialMedia);
    Render.append("contexts", contexts, ContextButton);
}

function onbind() {
    Render.clear("modpacks");
    Render.append("modpacks", view.data.modpacks, ModPack);
    document.getElementById("launch").onclick = view.launch;
}

function toggleSidebar(el) {
    if (el.getAttribute("data-sidebar") === "true") {
        el.setAttribute("data-sidebar", "false");
    } else {
        el.setAttribute("data-sidebar", "true");
    }
}

function action(name) {
    view.action(name);
}

function selectPack(el) {
    let all = document.getElementsByTagName(el.tagName);
    for (let i = 0; i < all.length; i++) {
        let e = all[i];
        e.setAttribute("data-selected", e === el);
    }
    view.selectPack(el.id);
}
