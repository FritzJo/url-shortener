function shrink_button_click(){
	var div = document.getElementById('loading-div');
	div.style.display = 'block';
    console.log("Processing...");
	console.log(shrink_url());
	div.style.display = 'none';
}

function shrink_url(url){
	//TODO
	var current_url = document.getElementById('icon_prefix').value
	return current_url
}
