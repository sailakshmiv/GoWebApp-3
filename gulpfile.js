var gulp = require('gulp');
var less = require('gulp-less');
var path = require('path');
var ts = require('gulp-typescript');

var lessPath = './less/**/*.less';
var typescriptPath = './typescript/**/*.ts'

gulp.task('less', function () {
  return gulp.src('less/site.less')
    .pipe(less({
      paths: [ path.join(__dirname, 'less', 'includes') ]
    }))
    .pipe(gulp.dest('./wwwroot/css'));
});

gulp.task('typescript', function() {
  console.log('Compiling typescript');
  return gulp.src([typescriptPath])
    .pipe(ts({module: 'commonjs'})).js.pipe(gulp.dest('./wwwroot/js'))
});

gulp.task('watch', function() {
  gulp.watch(lessPath, ['less']);
  gulp.watch(typescriptPath, ['typescript']);
});
