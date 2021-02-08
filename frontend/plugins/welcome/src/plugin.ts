import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import Login from '../../../packages/app/src/components/Login';
import Course from './components/CourseBuild';
import Courseclass from './components/Courseclass';
import EntSubjectsOffered from './components/SubjectsOffered'
import SearchInstructor from './components/SearchInstructor';
import searchsubjectsoffered from './components/SearchSubjectsOffered';
import FindCourse from './components/FindCourse';
import searchCourseclass from './components/searchCourseclass';
//import Sign from './components/Sign-in';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', WelcomePage);
    router.registerRoute('/welcome', WelcomePage);
    router.registerRoute('/course', Course);
    router.registerRoute('/courseclass', Courseclass);
    router.registerRoute('/subjectsoffered', EntSubjectsOffered);
    router.registerRoute('/SearchInstructor', SearchInstructor);
    router.registerRoute('/FindCourse', FindCourse);
    router.registerRoute('/searchsubjectsoffered', searchsubjectsoffered);
    router.registerRoute('/searchCourseclass', searchCourseclass);
    router.registerRoute('/Login', Login);
    //router.registerRoute('/Sign-in', Sign);
  },
});
