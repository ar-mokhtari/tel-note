$("#MainGrid").append("<div id=\"data-table\"></div>\n");
// $("#data-table").addClass("ag-theme-alpine");
$("#data-table").addClass("ag-theme-balham");
$("#data-table").addClass("w-100");
$("#data-table").css("height", "600px");

gridOptions = {

    defaultColDef: {
        sortable: true,
        filter: 'agTextColumnFilter',
        resizable: true
    },

    pagination: true,

    cacheQuickFilter: false,

    // columnDefs: columnDefs,
    onGridReady: (event) => {
        renderDataInTheTable(event.api, apiUrl)
    }
};

function onFilterTextBoxChanged() {
    gridOptions.api.setQuickFilter(
        document.getElementById('MainSearch').value
    );
}


eGridDiv = document.getElementById('data-table');
new agGrid.Grid(eGridDiv, gridOptions);


function renderDataInTheTable(api, apiUrl) {
    fetch(apiUrl)
        .then(function (response) {
            return response.json();
        }).then(function (data) {
        AllData = data
        // set the column headers from the data
        const colDefs = api.getColumnDefs();
        colDefs.length = 0;
        const keys = Object.keys((data.Data)[0])
        keys.forEach(key => colDefs.push({field: key}));
        api.setColumnDefs(colDefs);

        // add the data to the grid

        api.setRowData(data.Data);
        api.sizeColumnsToFit();
        $("#statusRespond").empty().append("<span> " + AllData.State + " </span>");

    });
}
