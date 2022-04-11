function initGrid(apiURL) {

    $("#MainGrid").append("<div id=\"data-table\"></div>\n");
    // $("#data-table").addClass("ag-theme-alpine");
    $("#data-table").addClass("ag-theme-balham");
    $("#data-table").addClass("w-100");
    $("#data-table").css("height", "600px");
    const gridOptions = {

        defaultColDef: {
            sortable: true,
            filter: 'agTextColumnFilter',
            resizable: true
        },

        pagination: true,

        // columnDefs: columnDefs,
        onGridReady: (event) => {
            renderDataInTheTable(event.api, apiURL)
        }
    };

    const eGridDiv = document.getElementById('data-table');

    new agGrid.Grid(eGridDiv, gridOptions);

}

function renderDataInTheTable(api, apiURL) {
    fetch(apiURL)
        .then(function (response) {
            return response.json();
        }).then(function (data) {
        // set the column headers from the data
        const colDefs = api.getColumnDefs();
        colDefs.length = 0;
        const keys = Object.keys((data.Data)[0])
        keys.forEach(key => colDefs.push({field: key}));
        api.setColumnDefs(colDefs);

        // add the data to the grid

        api.setRowData(data.Data);
        api.sizeColumnsToFit();
    });
}
