function shrinkButtonClick() {
    const resultDiv = document.getElementById("result-div");
    document.getElementById("text_label").classList.add("active");
    resultDiv.style.display = "block";
    const currentHostname = window.location.hostname;

    const apiurl = "http://" + currentHostname + ":8080/api/v1/short";
    // Get target url
    const params = document.getElementById("icon_prefix").value;
    const xhr = new XMLHttpRequest();

    xhr.onreadystatechange = function () {
        if (this.readyState !== 4) {
            return;
        }
        if (this.status === 200) {
            const data = this.responseText;
            const json = JSON.parse(data);
            document.getElementById("s-url").innerHTML = json.target;
            document.getElementById("t-url").innerHTML = json.source;

            //Show results
            document.getElementById("result-div").style.display = "block";
        }

    };
    xhr.open("POST", apiurl, true);
    xhr.send(params);
}
