const Render = {
    append: function(parent, source, comp) {
        if (typeof parent === "string") {
            parent = document.getElementById(parent);
        }

        if (source.forEach) {
            source.forEach(function(s) {
                Render.append(parent, s, comp);
            });
            return;
        } else {
            let component = new comp(source);
            let el = Render.render(component);
            parent.appendChild(el);
        }
    },

    render: function(comp) {
        if (typeof comp === "string") {
            return document.createTextNode(comp);
        }

        let el = document.createElement(comp.tag);
        Object.keys(comp).forEach(function(k) {
            switch (k) {
                case "tag":
                case "events":
                case "children":
                    break;
                default:
                    el.setAttribute(k, comp[k]);
            }
        });
        if (comp.events) {
            Object.keys(comp.events).forEach(function(k) {
                el.addEventListener(k, function(e) {
                    this.events[k].bind(this)(comp.el, e);
                }.bind(comp));
            });
        }
        if (comp.data) {
            Object.keys(comp.data).forEach(function(k) {
                el.setAttribute("data-" + k, comp.data[k]);
            });
        }
        if (comp.children && comp.children.forEach) {
            comp.children.forEach(function(child) {
                el.appendChild(Render.render(child));
            });
        }
        comp.el = el;
        return el;
    },

    clear: function(element) {
        if (typeof element === "string") {
            element = document.getElementById(element);
        }
        while (element.lastChild) {
            element.removeChild(element.lastChild);
        }
    },
};