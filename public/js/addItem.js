function setupIframeJS (e) {
    iframeDoc = e.srcElement.contentDocument
    iframeDoc.addEventListener("click", function (iframeClickEvent) {
        x = iframeClickEvent.clientX
        y = iframeClickEvent.clientY
        clickedEL = iframeDoc.elementFromPoint(x, y)
        console.log("----- value = ", clickedEL.innerText)
    })
}

function onItemAddSuccess (rsp) {  
    const previewURL = rsp.data.url

    const iframe = document.getElementById("preview")

    iframe.addEventListener("load", setupIframeJS)
    document.getElementById("preview").src=previewURL
}

function onItemAddFailure (err) {
    console.log(err)
}

function loadURL (e) {
    e.preventDefault()
    url = e.srcElement.elements[0].value
    if (!url.length) {
        alert("Please enter a URL")
        return
    }

    axios({
        method: 'post',
        url: '/api/items',
        data: { url }
    })
        .then(onItemAddSuccess)
        .catch(onItemAddFailure)
}