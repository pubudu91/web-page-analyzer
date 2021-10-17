const analyzeBtn = document.querySelector("#analyze")
const results = document.querySelector("#results")
const urlInput = document.querySelector("#url")

analyzeBtn.addEventListener("click", analyzeURL)
urlInput.addEventListener("input", validateURLInput)

// functions

function analyzeURL() {
    results.innerText = "Analyzing...."

    fetch("/analyze?" + new URLSearchParams({ url: urlInput.value }))
        .then(res => res.json())
        .then(data => {
            results.removeChild(results.firstChild)
            results.appendChild(createTable(data))
        })
        .catch(error => console.log(error))
}

function createTable(data) {
    let table = document.createElement("table")
    table.setAttribute("class", "paleBlueRows")

    // Status
    addRow(table, "Status:", data["Status"])

    // HTML version
    addRow(table, "HTML Version:", data["HtmlVersion"])

    // Title
    addRow(table, "Title:", data["Title"])

    // Headings
    let row = addRow(table, "Headings", "")
    row.firstChild.setAttribute("colspan", "2")
    row.removeChild(row.lastChild)

    for (const key in data["Headings"]) {
        const element = data["Headings"][key]
        row = addRow(table, key + ":", element)
        row.firstChild.setAttribute("style", "padding-left:5.00em;")
    }

    // Links
    row = addRow(table, "Links", "")
    row.firstChild.setAttribute("colspan", "2")
    row.removeChild(row.lastChild)

    for (const key in data["Links"]) {
        const element = data["Links"][key]
        row = addRow(table, key + ":", element)
        row.firstChild.setAttribute("style", "padding-left:5.00em;")
    }

    // Has Login
    addRow(table, "Has a log-in form:", data["HasLoginForm"])

    return table
}

function addRow(table, key, val) {
    let row = table.insertRow()

    let keyCell = row.insertCell();
    let text = document.createTextNode(key)
    keyCell.appendChild(text)

    let valCell = row.insertCell();
    text = document.createTextNode(val)
    valCell.appendChild(text)

    return row
}

function validateURLInput() {
    if (urlInput.value !== "" && !isValidURL(urlInput.value)) {
        document.querySelector("#errormsg").innerHTML = "Invalid HTTP URL!"
        return
    }

    document.querySelector("#errormsg").innerHTML = ""

    if (urlInput.value !== "") {
        analyzeBtn.disabled = false
    }
}

function isValidURL(givenURL) {
    let url;

    try {
        url = new URL(givenURL)
    } catch (_) {
        return false
    }

    return url.protocol === "http:" || url.protocol === "https:"
}
