const form = $("#form");

$(document).ready(function() {

    jQuery.validator.addMethod("username", function(value, element) {
        return this.optional(element) || /^[a-zA-Z\d]+$/.test(value);
    }, "Only alphanumeric characters");

    jQuery.validator.addMethod("password", function(value, element) {
        return this.optional(element) || /^[ -~]+$/.test(value);
    }, "Only printable characters");

    $("form").validate({
        rules: {
            username: {
                required: true,
                username: true,
                minlength: 6,
                maxlength: 30,
            },
            email: {
                required: true,
                email: true
            },
            password: {
                required: true,
                password: true,
                minlength: 6,
                maxlength: 30
            },
            password_confirm: {
                required: true,
                equalTo: "#password"
            }
        }
    })

    $(document).on('submit', form, function(e){
        e.preventDefault();

        $.ajax({
            type: "POST",
            url: "/sign_up",
            data: JSON.stringify({
                "username" : $('#username').val(),
                "email" : $('#email').val(),
                "password" : $('#password').val()}
            ),
            dataType: "json",
            contentType : "application/json",
            success: () => {
                location.href = "/confirm"
            },
            error: () => {
                location.href = "/sign_up"
            }
        });
    });
});
