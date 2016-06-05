$(function () {
    $('#contact-form').validator();
    $('#contact-form').on('submit', function (e) {
        if (!e.isDefaultPrevented()) {
	    var url = "http://localhost:8080/alumni"

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
		    var alertBox =
			'<div class="alert alert-success alert-dismissable"><button type="button" class="close" data-dismiss="alert" aria-hidden="true">&times;</button>Thank you! We\'ll be in touch!</div>';
		    $('#contact-form').find('.messages').html(alertBox);
		    $('#contact-form')[0].reset();
                }
            });

	    // prevent doing the regular submit shit
	    e.preventDefault();
        }
    })
});
