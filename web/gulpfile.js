var gulp = require('gulp')
/*, 
    uglify = require('gulp-uglify'); */

gulp.task('scripts', function() {
    return gulp.src('public/assets/**/*.js')
      .pipe(concat('main.js'))
      .pipe(gulp.dest('public/assets/js'))
      .pipe(rename({suffix: '.min'}))
      .pipe(uglify())
      .pipe(gulp.dest('dist/assets/js'))
      .pipe(notify({ message: 'Scripts task complete' }));
  });


  
exports.default = scripts