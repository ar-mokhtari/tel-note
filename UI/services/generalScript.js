const hostName = 'http://127.0.0.1:1212'
const getMenuUrl = hostName + '/menu-list';


let AllData;
let apiUrl;

jQuery.loadScript = function (url, callback) {
    jQuery.ajax({
        url: url,
        dataType: 'script',
        success: callback,
        async: true
    });
}

//set form active name in header
$("#headerResponse div a").text($("#mainTab li a.active").text());

//console.log(mask("5022291052143078","**** **** **** ****"))
function mask(value, pattern) {
    let count = 0;
    return pattern.replace(/\*/g, () => value[count++] || '')
}

//TODO::: handle this func for show Persian dates
function PersianData() {
    const date = new Date()
    const option = {
        weekday: "long",
        year: "numeric",
        month: "long",
        day: "numeric"
    }
    console.log(date.toLocaleDateString("fa-IR", option))
}

PersianData()

//menu click
$(document).on("click", "#MainSidebarAccordion div.menuItems", function () {
    let subMenu = $(this).text();
    let MenuPosition = $(this).parent().closest('div.accordion-item').find("h2").attr('id');
    let Menu = $("#" + MenuPosition).find("button").text();
    let formType = $(this).data('type');
    let route = hostName + $(this).data('route');
    // search input empty
    $("#MainSearch").val('');
    // form named
    $("#headerResponse div a").text(Menu);
    //tabs
    //menu types implement
    switch (formType) {
        //data entry from
        case 100:
            break;
        //list type from
        case 200:
            $("#MainGrid").empty();
            createTabs(subMenu);
            $.loadScript('UI/services/ag-girdPlacement.js', function () {
                apiUrl = route;
            });
            //native mode
            // createTabs(subMenu);
            // getData(route);
            // setTimeout("dbGridHtmlPlacement()", 50);
            break;
        //report type from
        case 500:
            $("#MainGrid").empty();
            createTabs(subMenu);
            $.loadScript('UI/services/ag-girdPlacement.js', function () {
                apiUrl = route;
            });
            //native mode
            // createTabs(subMenu);
            // getData(route);
            // setTimeout("dbGridHtmlPlacement()", 50);
            break;
        //action type from
        case 10000:
            break;
    }
});

//show tab management
// $("#mainTab a").click(function (e) {
//     e.preventDefault();
//     $("#MainSearch").val('');
//     $("#headerResponse div a").text($(this).text());
//     $("body tr").show();
//     $(this).tab("show");
// });

//grid search
$("#MainSearch").on("keyup", function () {
    let value = $(this).val().toLowerCase();
    let target = ".ag-row";
    if (value != "") {
        value = value.replace("ی", "ي");
        let isFindTr = false;
        $(target).each(function () {
            let targetTDs = $(this).find(".ag-cell")
            $(targetTDs).each(function () {
                if ($(this).text().replace("ی", "ي").toLowerCase().indexOf(value) != -1) {
                    isFindTr = true;
                    $(this).addClass("text-decoration-underline text-danger");
                    $(this).addClass("fw-bold");
                    $(this).addClass("fst-italic");
                } else {
                    $(this).removeClass("text-decoration-underline text-danger");
                    $(this).removeClass("fw-bold");
                    $(this).removeClass("fst-italic");
                }
            });
        });
    } else {
        $(".ag-row").show();
        $(".ag-cell").removeClass("text-decoration-underline text-danger");
        $(".ag-cell").removeClass("fw-bold");
        $(".ag-cell").removeClass("fst-italic");
    }
});

// first create tab structure
function createTabs(subMenu) {
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
