function menuPlacement(data) {
    $.each(data, function (menuIndex, menuElement) {
        $("#MainSidebar").append(
            "<li class='mb-1'>" +
            "<button aria-expanded=\"false\" class=\"btn btn-toggle align-items-center rounded collapsed\"" +
            " data-bs-target = #group-collapse" + menuIndex + " data-bs-toggle = collapse > " +
            menuElement.GroupName +
            "</button>" +
            "<div class=\"collapse\" id=\"" + "group-collapse" + menuIndex + "\">\n" +
            "    <ul class=\"btn-toggle-nav list-unstyled fw-normal pb-1 small\">\n" +
            "    </ul>\n" +
            "</div>" +
            "</li>"
        );
        $.each(menuElement.Row, function (detailIndex, detailElement) {
            $("#group-collapse" + menuIndex + " ul").append("<li><a class=\"link-dark rounded\" href=\"#\">" + detailElement.Description + "</a></li>")
        })
    })
}
