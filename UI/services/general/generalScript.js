//set form active name in header
$("#headerResponse div a").text($("#mainTab li a.active").text());

//console.log(mask("5022291052143078","**** **** **** ****"))
function mask(value, pattern) {
    let count = 0;
    return pattern.replace(/\*/g, () => value[count++] || '')
}

$(document).on("click", "#MainSidebar a", function () {
    let subMenu = $(this).text();
    let MenuPosition = $(this).parent().closest('div').attr('id');
    let Menu = $("#" + MenuPosition).closest("li").find("button").text();
    // search input empty
    $("#MainSearch").val('');
    // form named
    $("#headerResponse div a").text(Menu);
    // $("body tr").show();
    createGrids(subMenu);
});

$("#mainTab a").click(function (e) {
    e.preventDefault();
    $("#MainSearch").val('');
    $("#headerResponse div a").text($(this).text());
    $("body tr").show();
    $(this).tab("show");
});

//offline search
$("#MainSearch").on("keyup", function () {
    let value = $(this).val().toLowerCase();
    value = value.replace("ی", "ي");
    let target = "#tab-content div.active div table tbody tr";
    let isFindTr = false;
    $(target).each(function () {
        let targetTDs = $(this).find("td")
        $(targetTDs).each(function () {
            if (value != "" && $(this).text().replace("ی", "ي").toLowerCase().indexOf(value) !== -1) {
                isFindTr = true;
                $(this).addClass("text-danger");
            } else {
                $(this).removeClass("text-danger");
            }
        });
        alert($(this).find("td:last-child").text());
        switch (isFindTr) {
            case true:
                $(this).eq($(this).index() + 1).show();
            case false:
                $(this).eq($(this).index() + 1).hide();
        }
        isFindTr = false;
    });
});
