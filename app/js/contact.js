$(function () {
    $('#contact-form').validator();
    $('#contact-form').on('submit', function (e) {
        if (!e.isDefaultPrevented()) {
	    var url = "http://roll-tech.appspot.com/alumni"

	    var data = {};
	    $.each(this.elements, function(i, v){
		var input = $(v);

		// This is gross but fuck you
		if (input.attr("name") == "year") {
		    data[input.attr("name")] = Number(input.val());
		} else {
		    data[input.attr("name")] = input.val();
		}
		delete data["undefined"];
	    });

            $.ajax({
                type: "POST",
                url: url,
                data: JSON.stringify(data),
		dataType: "json",
                success: function (data){
		    console.log(JSON.parse(data));
		    $(this).html("Success!");
                }
            });

	    // prevent doing the regular submit shit
	    e.preventDefault();
        }
    })
});
