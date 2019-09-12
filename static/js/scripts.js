function shrink_button_click() {
    //Show loading bar
    var div = document.getElementById('loading-div');
    div.style.display = 'block';

    var current_url = document.getElementById('icon_prefix').value
    var parameter = window.location.href
    window.location = "/index.html?target=" + current_url;

    //Hide loading bar
    div.style.display = 'none';

    //Show results
    div = document.getElementById('result-div');
    div.style.display = 'block';
}

window.onload = function () {
    var parameter = window.location.href
    if (parameter.includes("target")) {
        var url = new URL(parameter);
        //document.addEventListener('turbolinks:load', () => {
        //    M.updateTextFields();
        //});
        document.getElementById("text_label").classList.add('active');
        document.getElementById('icon_prefix').value = url.searchParams.get("target");
        div = document.getElementById('result-div');
        div.style.display = 'block';
    }
};
