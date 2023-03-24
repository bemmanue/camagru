$(document).ready(function() {

    $(".like_button").click(function(e) {
        e.preventDefault();

       console.log(this.closest(".post").id)

        $.ajax({
            type: "POST",
            url: "/like",
            data: JSON.stringify({
                "post_id" : parseInt(this.closest(".post").id)
            }),
            dataType: "json",
            contentType : "application/json",
            success: () => {
                const btn = this.querySelector(".like_svg")

                const likes = this.parentNode.querySelector(".like_count")
                console.log(likes.innerHTML)
                if (btn.classList.contains("dislike_img")) {
                    likes.innerHTML = (parseInt(likes.innerHTML) + 1).toString()
                } else {
                    likes.innerHTML = (parseInt(likes.innerHTML) - 1).toString()
                }

                btn.classList.toggle("dislike_img")
                btn.classList.toggle("like_img")
            },
            error: () => {

            }
        });
    });
});