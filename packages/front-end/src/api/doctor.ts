import { api } from "./api";
import type { Log } from "@/type/log";
import type { Patient, Res, Bed } from "@/type";
import type {AddPatientRequest} from "@/type/request";

export function DoctorGetPatients() {
  return api.get<Res<Patient[]>>("/doctor/patients");
}

export function DoctorDeletePatient(id: number) {
  return api.delete<Res<null>>("/doctor/patient/" + id);
}

export function DoctorAddPatient(patient: AddPatientRequest) {
  return api.post<Res<null>>("/doctor/patient", patient);
}

export function DoctorGetBeds() {
  return api.get<Res<Bed[]>>("/doctor/beds");
}

export function DoctorGetPatient(id: number) {
  return api.get<Res<Patient>>("/doctor/patient/" + id);
}

export function DoctorAddLog(patient_id: number, content: string) {
  return api.post<Res<null>>("/doctor/log", { patient_id, content });
}

export function DoctorGetAllLogs() {
  return api.get<Res<Log[]>>("/doctor/logs");
}

export function DoctorGetLogs(patient_id: number) {
  return api.get<Res<Log[]>>("/doctor/log/" + patient_id);
}

export function DoctorGetPatientLogs(patient_id: number) {
  return api.get<Res<Log[]>>("/doctor/logs/"+patient_id.toString());
}

// export function DoctorGetVitalSigns()
