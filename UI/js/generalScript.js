//console.log(mask("5022291052143078","**** **** **** ****"))
function mask(value, pattern) {
    let count = 0;
    return pattern.replace(/\*/g, () => value[count++] || '')
}

$("#mainTab a").click(function (e) {
    e.preventDefault();
    $(this).tab("show");
});

$("#MainSearch").on("keyup", function () {
    let value = $(this).val().toLowerCase();
    value = value.replace("ی", "ي")
    let target = "#ResponseTabContent div table tbody tr"
    $(target).each(function () {
        let targetString = $(this).find("td:first").text();
        if (
            targetString.replace("ی", "ي").toLowerCase().indexOf(value) !== -1) {
            $(this).show();
        } else {
            $(this).hide();
        }
    });
});
