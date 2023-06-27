/*背景滾動*/
window.addEventListener('scroll', function() {
    var introSection = document.querySelector('.intro-section');
    var scrollPosition = window.scrollY;

    introSection.style.transform = 'scale(' + (1 + scrollPosition * 0.001) + ')';
}); 