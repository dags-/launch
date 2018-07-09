let anchor = null;

addEventListener("resize", function(e) {
    if (anchor !== null) {
        document.getElementById("content").scrollTop = anchor.offsetTop;
    }
});

function autoScroll(container, to, duration) {
    if (duration <= 0) {
        return;
    }
    let difference = to.offsetTop - container.scrollTop;
    let perTick = difference / duration * 10;
    setTimeout(function () {
        container.scrollTop += perTick;
        if (container.scrollTop === to.offsetTop) {
            anchor = to;
            return;
        }
        autoScroll(container, to, duration - 10);
    }, 10);
}