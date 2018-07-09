addEventListener("load", function() {
    document.body.addEventListener("keypress", reloadListener);
    window.addEventListener("resize", resizeListener);
    requestView();
});

function reloadListener(e) {
    if (e.ctrlKey && (e.key === "r" || e.key === "R")) {
        window.location.reload(true);
    }
}

function resizeListener(e) {
    if (view && view.setSize !== undefined) {
        view.setSize(window.innerWidth, window.innerHeight)
    }
}

function requestView() {
    let view = window.location.pathname;
    let get = new XMLHttpRequest();
    get.open("GET", "/view" + view, true);
    get.send(null);
}