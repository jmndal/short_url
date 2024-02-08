function uploadCSV() {
  const fileInput = document.getElementById("csvFile");
  const file = fileInput.files[0];

  if (file) {
    const formData = new FormData();
    formData.append("file", file);

    fetch("/upload", {
      method: "POST",
      body: formData,
    })
      .then((response) => response.json())
      .then((data) => displayData(data))
      .catch((error) => console.error("Error:", error));
  } else {
    console.error("Please select a file.");
  }
}

function displayData(data) {
  const table = document.getElementById("dataTable");
  table.innerHTML = "";

  // Add table headers
  const headers = Object.keys(data[0]);
  const headerRow = table.insertRow();
  headers.forEach((header) => {
    const th = document.createElement("th");
    th.textContent = header;
    headerRow.appendChild(th);
  });

  // Add table rows
  data.forEach((rowData) => {
    const row = table.insertRow();
    headers.forEach((header) => {
      const cell = row.insertCell();
      cell.textContent = rowData[header];
    });
  });
}
