function shrink_button_click() {
    //Show loading bar
    var div = document.getElementById('loading-div');
    div.style.display = 'block';

    //calculate and set results
    var current_url = document.getElementById('icon_prefix').value
    //var s_url = shrink_url(current_url)
    //document.getElementById('s-url').innerHTML = s_url;
    document.getElementById('t-url').innerHTML = current_url;
    //document.getElementById('date-url').innerHTML = "2020-02-20";
    //Hide loading bar
    div.style.display = 'none';

    //Show results
    div = document.getElementById('result-div');
    div.style.display = 'block';
}

function shrink_url(url) {
    //TODO
    return url
}
