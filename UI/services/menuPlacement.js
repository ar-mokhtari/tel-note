//first call menu to get data from server
function getMenu() {
    $("#MainSidebarAccordion").empty();
    $.ajax({
        type: 'GET',
        url: getMenuUrl,
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        data: JSON.stringify({GroupName: "GroupName", Row: "Row"}),
        Row: JSON.stringify({
            Title: "Title",
            Description: "Description",
            Route: "Route",
            Type: "Type",
        }),
        success: function (data) {
            menuPlacement(data);
        }
    });
}

//then placement in html
function menuPlacement(data) {
    $.each(data, function (menuIndex, menuElement) {
        $("#MainSidebarAccordion").append(
            "  <div class=\"accordion-item\">\n" +
            "    <h2 class=\"accordion-header\" id=\"" + "heading" + menuIndex + "\">\n" +
            "<button class=\"accordion-button collapsed\" type=\"button\" data-bs-toggle=\"" + "collapse" + "\" data-bs-target=\"" + "#collapse" + menuIndex + "\"" + " aria-expanded=\"false\" aria-controls=\"" + "collapse" + menuIndex + "\">\n" +
            menuElement.GroupName +
            "      </button>\n" +
            "    </h2>\n" +
            "    <div id=\"" + "collapse" + menuIndex + "\" class=\"accordion-collapse collapse\" aria-labelledby=\"" + "heading" + menuIndex + "\"  data-bs-parent=\"#MainSidebarAccordion\">\n" +
            "      <div class=\"accordion-body\">\n" +
            "        " +
            "      </div>\n" +
            "    </div>\n" +
            "  </div>"
        );
        $.each(menuElement.Row, function (detailIndex, detailElement) {
            $("#collapse" + menuIndex + " div.accordion-body").append(
                "<div class='link-dark text-cente  pb-1 small'><div style='cursor: pointer'  data-route=\"" + detailElement.Route + "\"  data-type=\"" + detailElement.Type + "\" class=\"menuItems rounded\" \"\">" +
                detailElement.Description +
                "</div></a>");
        });
    });
}
