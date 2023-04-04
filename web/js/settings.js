$(document).ready(function() {
    let comment_notify = document.getElementById("comment-notify")
    let like_notify = document.getElementById("like-notify")

    comment_notify.addEventListener("change", function () {
        let value = $(comment_notify).is(':checked')

        $.ajax({
            type: "POST",
            url: "/settings",
            data: JSON.stringify({
                "comment_notify" : value,
            }),
            dataType: "json",
            contentType : "application/json",
            success: function() {
            },
            error: function() {
            },
        })
    })

    like_notify.addEventListener("change", function () {
        let value = $(like_notify).is(':checked')

        $.ajax({
            type: "POST",
            url: "/settings",
            data: JSON.stringify({
                "like_notify" : value,
            }),
            dataType: "json",
            contentType : "application/json",
            success: function() {
            },
            error: function() {
            },
        })
    })
})