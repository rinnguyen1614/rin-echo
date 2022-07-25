export interface AcademicStats {
  id: number;
  title: string;
  count: string;
  new: string;
  badgeColor: string;
  bgcolor: string;
  icon: string;
}

export interface CourseCategory {
  id: number;
  title: string;
  desc: string;
  lessons: number;
  xp: number;
  images: {
    image: string;
    title: string;
  }[];
}

export interface Profile {
  id: number;
  profile_pic: string;
  name: string;
  designation: string;
  achievements: number;
  friends: number;
}

export interface Courses {
  categories: {
    id: number;
    title: string;
    slug: string;
  }[];
  courses: {
    id: number;
    title: string;
    duration: string;
    rating: number;
    isCompleted: boolean;
    thumb: string;
  }[];
}

export interface NotificationData {
  id: number;
  bgcolor: string;
  color: string;
  letter: string;
  content: string;
  date: any;
}

export interface CourseDetails {
  id: number;
  title: string;
  thumb: string;
  level: string;
  coveredDuration: string;
  totalDuration: string;
  coveredPractice: string;
  totalPractice: string;
  graphData: {
    month: string;
    duration: number;
  }[];
}

export interface LearningData {
  id: number;
  icon: string;
  title: string;
  desc: string;
  percentage: number;
}

export interface LatestResultData {
  id: number;
  chapter: string;
  topic: string;
  percentage: number;
}

export interface ClassData {
  id: number;
  name: string;
  percent: number;
  icon: string;
}

export interface StudentRankingData {
  id: number;
  name: string;
  profile_pic: string;
  courseId: number;
  courseName: string;
  totalGrade: number;
  ranking: number;
  category: string;
}

export interface Grades {
  month: string;
  grades: number;
}

export interface RelatedCoursesData {
  id: number;
  image: string;
  title: string;
  author: string;
  views: string;
}

export interface VideoPromoData {
  title: string;
  owner: string;
  category: string;
  assignments: {
    id: number;
    title: string;

    desc: string;
    students: number;
    daysLeft: number;
  }[];
}

export interface Academy {
  academicStats: AcademicStats[];
  courseCategories: CourseCategory[];
  profile: Profile;
  courses: Courses;
  notifications: NotificationData[];
  courseDetails: CourseDetails[];
  learningData: LearningData[];
  latestResults: LatestResultData[];
  classData: ClassData[];
  studentRankings: StudentRankingData[];
  grades: Grades[];
  relatedCourses: RelatedCoursesData[];
  videoPromo: VideoPromoData;
}
