const socials = [
    {name: "Discord", id: "discord", url: "https://google.com", icon: "mdi-discord"},
    {name: "Github", id: "github", url: "https://github.com/Conquest-Reforged", icon: "mdi-github-box"},
    {name: "Patreon", id: "patreon", url: "https://google.com", icon: "mdi-patreon"},
    {name: "Twitter", id: "twitter", url: "https://twitter.com/ConReforged", icon: "mdi-twitter"},
];

const contexts = [
    {name: "Settings", icon: "mdi-settings", action: "/settings"},
    {name: "Mods", icon: "mdi-format-list-checks", action: "settings"},
    {name: "Game Folder", icon: "mdi-folder-outline", action: "documents"},
    {name: "Screenshots", icon: "mdi-image-area", action: "screenshots"},
    {name: "Saves", icon: "mdi-earth", action: "saves"},
    {name: "Logs", icon: "mdi-console", action: "logs"},
    {name: "Crash Reports", icon: "mdi-file-outline", action: "reports"},
    {name: "Delete", icon: "mdi-delete", action: "delete"},
];

function SocialMedia(s) {
    this.tag = "a";
    this.class = "social-button " + s.id;
    this.href = s.url;
    this.target = "_blank";
    this.children = [
        {tag: "i", class: "mdi " + s.icon},
        {tag: "span", children: [s.name]}
    ];
}

function ModPack(pack) {
    this.tag = "div";
    this.class = "modpack";
    this.id = pack.id;
    this.onclick = "selectPack(this)";
    this.data = {selected: pack.selected};
    this.children = [
        {tag: "div", class: "title", children: [pack.name]},
        {tag: "div", class: "version", children: ["Version ", pack.version]},
    ];
}

function ContextButton(c) {
    this.tag = "i";
    this.title = c.name;
    this.onclick = "view.action('" + c.action + "')";
    this.class = "mdi " + c.icon + " context button hover";
}