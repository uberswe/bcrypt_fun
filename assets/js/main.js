
function generateBcryptHashes() {
    const difficulty = $("#difficulty").val();
    if (difficulty > 20) {
        alert("Cost is greater than 20, the cost will be reduced to 20.")
    }
    const values = $ ("#bcryptform").serializeArray ();
    const data = {};

    for (let i = 0; i < values.length; i++) {
        data[values[i].name] = values[i].value;
    }

    $("#loader").show()

    $.ajax({
        method: "POST",
        url: "/api/v1/hashes",
        data: data
    }).done(function (obj) {
        $("#loader").hide()
        const passblock = $ ("#password-block");
        let passblockstring = "";
        passblock.html("");
        for (i = 0; i < obj.hashes.length; i++) {
            if (i > 0) {
                passblockstring += "\n"
                passblockstring += obj.hashes[i].hash
            } else {
                passblockstring = obj.hashes[i].hash
            }
        }


        passblock.val(passblockstring);

        if (data["highlight"]) {
            selectTextareaLine(passblock, 1);
        }
    });
    return false;
}

// See https://stackoverflow.com/a/13651036/1260548
function selectTextareaLine(tarea,lineNum) {
    lineNum--; // array starts at 0
    const lines = tarea.value.split ("\n");

    // calculate start/end
    let startPos = 0;
    for(let x = 0; x < lines.length; x++) {
        if(x === lineNum) {
            break;
        }
        startPos += (lines[x].length+1);

    }

    var endPos = lines[lineNum].length+startPos;

    // do selection
    // Chrome / Firefox

    if(typeof(tarea.selectionStart) !== "undefined") {
        tarea.focus();
        tarea.selectionStart = startPos;
        tarea.selectionEnd = endPos;
        return true;
    }

    // IE
    if (document.selection && document.selection.createRange) {
        tarea.focus();
        tarea.select();
        const range = document.selection.createRange ();
        range.collapse(true);
        range.moveEnd("character", endPos);
        range.moveStart("character", startPos);
        range.select();
        return true;
    }

    return false;
}

generateBcryptHashes();
