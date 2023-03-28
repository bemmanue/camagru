const form = $("#form");

$(document).ready(function() {

    $(document).on('submit', form, function(e){
        e.preventDefault();
        $.ajax({
            type: "POST",
            url: "/sign_in",
            data: JSON.stringify({
                "username" : $('#username').val(),
                "password" : $('#password').val(),
            }),
            dataType: "json",
            contentType : "application/json",
            success: function() {
                location.replace("/feed")
                localStorage.setItem("username", $('#username').val())
            },
            error: () => {
                window.location.replace("/sign_in")
            },
        });
    });
});
