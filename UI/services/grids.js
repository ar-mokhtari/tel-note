function createGrids(subMenu) {
    alert();
    $("#MainGrid").empty();
    $("#MainGrid").append(
        "       <!-- Tabs -->\n" +
        "        <ul id=\"mainTab\" class=\"nav nav-tabs\">\n" +
        "            <li class=\"nav-item\"><a class=\"nav-link active\" href=\"#ResponseTabContent\">" + subMenu + "</a></li>\n" +
        "        </ul>\n" +
        "        <!-- Tabs -->\n" +
        "\n" +
        "        <!-- Response body -->\n" +
        "        <div id=\"tab-content\" class=\"tab-content\">\n" +
        "            <div class=\"tab-pane fade show active\" id=\"ResponseTabContent\">\n" +
        "                <div></div>\n" +
        "            </div>\n" +
        "        </div>\n" +
        "        <!-- Response body -->\n");
}
