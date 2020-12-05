function getSelector(node) {
    var id = node.getAttribute('id');
    if (id) {
        return '#'+id;
    }
    var path = '';
    while (node) {
        var name = node.localName;
        var parent = node.parentNode;
        if (!parent) {
            path = name + ' > ' + path;
            continue;
        }
        if (node.getAttribute('id')) {
            path = '#' + node.getAttribute('id') + ' > ' + path;
            break;
        }
        var sameTagSiblings = [];
        var children = parent.childNodes;
        children = Array.prototype.slice.call(children);
        children.forEach(function(child) {
            if (child.localName == name) {
                sameTagSiblings.push(child);
            }
        });
        // if there are more than one children of that type use nth-of-type
        if (sameTagSiblings.length > 1) {
            var index = sameTagSiblings.indexOf(node);
            name += ':nth-of-type(' + (index + 1) + ')';
        }
        if (path) {
            path = name + ' > ' + path;
        } else {
            path = name;
        }
        node = parent;
    }
    return path;
}

function setupIframeJS (e) {
    iframeDoc = e.srcElement.contentDocument
    iframeDoc.addEventListener("click", function (iframeClickEvent) {
        x = iframeClickEvent.clientX
        y = iframeClickEvent.clientY
        clickedEL = iframeDoc.elementFromPoint(x, y)
        console.log("----- value = ", clickedEL.innerText, getSelector(clickedEL))
    })
}

function onItemAddSuccess (rsp) {  
    const previewURL = rsp.data.cachedURL
    const urlID = rsp.data.urlID
    const iframe = document.getElementById("preview")
    iframe.addEventListener("load", setupIframeJS)
    iframe.getAttribute("data-urlID", urlID)
    iframe.src=previewURL
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
        url: '/urls',
        data: { url }
    })
        .then(onItemAddSuccess)
        .catch(onItemAddFailure)
}
