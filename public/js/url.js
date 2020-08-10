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

function clickedIframeElement (e) { 
    x = e.clientX
    y = e.clientY
    iframeDoc = document.querySelector("#preview").contentDocument
    clickedEL = iframeDoc.elementFromPoint(x, y)

    currentValue = clickedEL.innerText
    selector = getSelector(clickedEL)

    instr = document.querySelector("#instruction")
    if (instr) {
        instr.remove()
    }

    newTracker = document.querySelector("#tracker-template").cloneNode(true)
    newTracker.removeAttribute("id")
    newTracker.querySelector(".tracker-value").setAttribute("value", currentValue)
    newTracker.querySelector(".tracker-selector").setAttribute("value", selector)
    document.querySelector("#tracker-form").append(newTracker)
    newTracker.classList.remove("hidden")
}

function setupIframeJS (e) {
    iframeDoc = e.srcElement.contentDocument
    iframeDoc.addEventListener("click", clickedIframeElement)
}

document.querySelector("#preview").addEventListener("load", setupIframeJS)

function addTrackers (e) {
    e.preventDefault()
    const form = new FormData(e.srcElement)
    const urlID = form.get("urlID")
    const names = form.getAll("tracker-name")
    const selectors = form.getAll("tracker-selector")

    const trackers = names.map(function(name, i) {
        return {
            name: name,
            selector: selectors[i]
        }
    })

    axios({
        method: 'post',
        url: `/urls/${urlID}/trackers`,
        data: {
            trackers: trackers
        }
    })
        .then(console.log.bind(console, "added"))
        .catch(console.log.bind(console, "error"))
}   
