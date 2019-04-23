// $('.nav li').on('click', function(){
//     $('.nav li').removeClass('active');
//     $(this).addClass('active');
// });

// $(document).ready(function() {
//     $('.nav li').removeClass('active');
//     // get current URL path and assign 'active' class
//     var pathname = window.location.pathname;
//     $('.nav > li > a[href="'+pathname+'"]').parent().addClass('active').append(" <span class='sr-only'>(current)></span>")
// })

// $('.navbar-nav .li').click(function(){
//     $('.navbar-nav .li').removeClass('active');
//     $(this).addClass('active');
// });

$( '#topheader .navbar-nav a' ).on( 'click', function () {
    $( '#topheader .navbar-nav' ).find( 'li.active' ).removeClass( 'active' );
    $( this ).parent( 'li' ).addClass( 'active' );
});

