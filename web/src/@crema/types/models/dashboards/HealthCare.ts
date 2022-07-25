export interface DrState {
  id: number;
  category: string;
  name: string;
  bgColor: string;
  time: string;
  icon: string;
}

export interface HeathStaticsData {
  dataOne: {
    month: string;
    number: number;
  }[];
  dataTwo: {
    month: string;
    number: number;
  }[];
  dataThree: {
    month: string;
    number: number;
  }[];
}

export interface NewpatientsData {
  month: string;
  number: number;
}

export interface CancelVisitData {
  month: string;
  number: number;
}

export interface TopDoctorData {
  id: number;
  name: string;
  speciality: string;
  profile_pic: string;
  rating: string;
  scheduled: boolean;
}

export interface UpcomingAppointment {
  id: number;
  name: string;
  speciality: string;
  profile_pic: string;
  appointmentTime: string;
  appointmentDate: string;
}

export interface NotificationData {
  id: number;
  title: string;
  time: string;
  color: string;
}

export interface RecentPatientData {
  id: string;
  name: string;
  profile_pic: string;
  gender: string;
  weight: string;
  assignedDr: string;
  date: string;
  status: string;
  color: string;
}

export interface HospitalActivityData {
  name: string;
  consultations: number;
  patients: number;
}

export interface BloodCard {
  id: number;
  name: string;
  icon: string;
  measurement: string;
  color: string;
}

export interface AppointmentCards {
  id: number;
  name: string;
  value: string;
  icon: string;
  chartValue: string;
  chartTime: string;
  chartData: {
    month: string;
    users: number;
  }[];
  color: string;
}

export interface HeartCard {
  id: number;
  title: string;
  measurement: string;
  unit: string;
  graphData: {
    name: string;
    rate: number;
  }[];
  color: string;
}

export interface Doses {
  id: number;
  value: string;
  name: string;
  icon: string;
  bgColor?: string;
}

export interface YourActivityData {
  day: string;
  visits: number;
}

export interface HealthCare {
  drState: DrState[];
  heathStatics: HeathStaticsData;
  newpatients: NewpatientsData[];
  cancelVisits: CancelVisitData[];
  topDoctors: TopDoctorData[];
  upcomingAppointment: UpcomingAppointment[];
  notifications: NotificationData[];
  hospitalStatics: Doses[];
  recentpatients: RecentPatientData[];
  hospitalActivity: HospitalActivityData[];
  bloodCard: BloodCard[];
  appointmentCards: AppointmentCards[];
  heartCard: HeartCard;
  temperatureCard: HeartCard;
  doses: Doses[];
  yourActivity: YourActivityData[];
}
