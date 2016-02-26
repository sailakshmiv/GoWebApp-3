var gulp = require('gulp');
var less = require('gulp-less');
var path = require('path');

var lessPath = './less/**/*.less';

gulp.task('less', function () {
  return gulp.src('less/site.less')
    .pipe(less({
      paths: [ path.join(__dirname, 'less', 'includes') ]
    }))
    .pipe(gulp.dest('./wwwroot/css'));
});

gulp.task('watch', function() {
  gulp.watch(lessPath, ['less']);
});
