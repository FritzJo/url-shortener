function shrinkButtonClick() {
    //Show loading bar
    var div = document.getElementById("loading-div");
    div.style.display = "block";

    var currentUrl = document.getElementById("icon_prefix").value;
    var parameter = window.location.href;
    window.location = "/index.html?target=" + currentUrl;

    //Hide loading bar
    div.style.display = "none";

    //Show results
    div = document.getElementById("result-div");
    div.style.display = "block";
}

window.onload = function () {
    var parameter = window.location.href;
    if (parameter.includes("target")) {
        const url = new URL(parameter);
        const resultDiv = document.getElementById("result-div");
        document.getElementById("text_label").classList.add("active");
        document.getElementById("icon_prefix").value = url.searchParams.get("target");
        resultDiv.style.display = "block";
    }
};