function shrinkButtonClick() {
    //Show loading bar
    var div = document.getElementById("loading-div");
    div.style.display = "block";

    const resultDiv = document.getElementById("result-div");
    document.getElementById("text_label").classList.add("active");
    resultDiv.style.display = "block";
    var currentHostname = window.location.hostname

    var apiurl = "http://" + currentHostname + ":8080/api/v1/short";
    // Get target url
    var params = document.getElementById("icon_prefix").value;
    var xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function () {
        if (this.readyState != 4) return;
        if (this.status == 200) {
            var data = this.responseText;
            var json = JSON.parse(data);
            document.getElementById("s-url").innerHTML = json.target;
            document.getElementById("t-url").innerHTML = json.source;

            //Hide loading bar
            div.style.display = "none";
            //Show results
            div = document.getElementById("result-div");
            div.style.display = "block";
        }

    };
    xhr.open("POST", apiurl, true);
    xhr.send(params);
}
