const analyzeBtn = document.querySelector("#analyze")
const results = document.querySelector("#results")
const urlInput = document.querySelector("#url")
const diagnosticMsg = document.querySelector("#errorMsg")

analyzeBtn.addEventListener("click", analyzeURL)
urlInput.addEventListener("input", validateURLInput)

// functions

function analyzeURL() {
    diagnosticMsg.innerHTML = "Analyzing..."
    diagnosticMsg.setAttribute("class", "normalDiagnostic")

    if (results.firstChild !== null) {
        results.removeChild(results.firstChild)
    }

    invokeAnalyzeEndpoint(urlInput.value)
}

async function invokeAnalyzeEndpoint(url) {
    try {
        let response = await fetch("/analyze?" + new URLSearchParams({ url: url }))
        let data = await response.json()

        if (data == null) {
            setErrorDiagnostic("[Status: " + response.status +
                "] Server failed to process the web page returned by the specified URL.")
            return
        }

        if (data.hasOwnProperty("error")) {
            setErrorDiagnostic("[Status: " + response.status +
                "] Failed to anaylze the web page. Cause: " + data["error"])
            return
        }

        diagnosticMsg.innerHTML = ""
        results.appendChild(createTable(data))
    } catch (err) {
        setErrorDiagnostic("Failed to fetch the analysis. Please try again.")
        console.log(err)
    }
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
        row.firstChild.setAttribute("style", "padding-left: 75px;")
    }

    // Links
    row = addRow(table, "Links", "")
    row.firstChild.setAttribute("colspan", "2")
    row.removeChild(row.lastChild)

    for (const key in data["Links"]) {
        const element = data["Links"][key]
        row = addRow(table, key + ":", element)
        row.firstChild.setAttribute("style", "padding-left: 75px;")
    }

    // Has Login
    addRow(table, "Has a log-in form:", data["HasLoginForm"])

    return table
}

function addRow(table, key, val) {
    let row = table.insertRow()

    let keyCell = row.insertCell()
    let text = document.createTextNode(key)
    keyCell.appendChild(text)

    let valCell = row.insertCell()
    text = document.createTextNode(val)
    valCell.appendChild(text)

    return row
}

function validateURLInput() {
    if (urlInput.value !== "" && !isValidURL(urlInput.value)) {
        setErrorDiagnostic("Invalid HTTP URL!")
        analyzeBtn.disabled = true
        return
    }

    diagnosticMsg.innerHTML = ""

    if (urlInput.value !== "") {
        analyzeBtn.disabled = false
    }
}

function isValidURL(givenURL) {
    let url

    try {
        url = new URL(givenURL)
    } catch (_) {
        return false
    }

    return url.protocol === "http:" || url.protocol === "https:"
}

function setErrorDiagnostic(msg) {
    diagnosticMsg.innerHTML = msg
    diagnosticMsg.setAttribute("class", "errorDiagnostic")
}
