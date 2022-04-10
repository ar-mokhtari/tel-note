$('head').append('<link rel="stylesheet" href="../SDK/ag-grid/ag-grid.css" type="text/css" />');
$('head').append('<link rel="stylesheet" href="../SDK/ag-grid/ag-theme-balham.css" type="text/css" />');


$("#MainGrid").empty();
$("#MainGrid").append("<style>\n" +
    "    #data-table {\n" +
    "        height: 500px;\n" +
    "        width: 100%;\n" +
    "    }\n" +
    "</style>\n" +
    "\n" +
    "\n" +
    "<div id=\"data-table\" class=\"ag-theme-balham\">\n" +
    "</div>\n");
const columnDefs = [
    {field: "Id"},
    {field: "Name"},
    {field: "EnglishName"},
    {field: "AriaCode"},
    {field: "Lat"},
    {field: "Lng"},
];

const gridOptions = {

    defaultColDef: {
        sortable: true,
        filter: 'agTextColumnFilter',
        resizable: true
    },

    pagination: true,

    columnDefs: columnDefs,
    onGridReady: (event) => {
        renderDataInTheTable(event.api)
    }
};

const eGridDiv = document.getElementById('data-table');

new agGrid.Grid(eGridDiv, gridOptions);

function renderDataInTheTable(api) {
    fetch("http://127.0.0.1:1212/get-city")
        .then(function (response) {
            return response.json();
        }).then(function (data) {
        api.setRowData(data.Data);
        api.sizeColumnsToFit();
    })
}
