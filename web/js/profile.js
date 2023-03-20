const form = $("#form");

$(document).ready(function() {

    $(document).on('submit', form, function(e){

        e.preventDefault();
        $.ajax({
            type: "POST",
            url: "/upload",
            data: new FormData(document.getElementById("form")),
            contentType: false,
            processData: false,
            success: function() {
                location.replace("/profile")
            },
            error: () => {
                location.replace("/profile")
            },
        });
    });
});
