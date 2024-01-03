import {api} from "./api";
import type { Bed,
  Department,
  Doctor,
  Patient,
  Res,
  PatientDetail,
  VitalSign
} from "@/type";

export function AdminGetDepartments(){
  return api.get<Res<Department[]>>("/admin/department");
}

export function AdminAddDepartment(name:string){
  return api.post<Res<null>>("/admin/department",{name});
}

export function AdminUpdateDepartment(
  id: number,
  name: string
) {
  return api.put<Res<null>>("/admin/department", {
    id, name
  });
}

export function AdminDeleteDepartment(id: number) {
  return api.delete<Res<null>>("/admin/department",{
    params: { id }
  });
}

export function AdminGetAllDoctors(){
  return api.get<Res<Doctor[]>>("/admin/doctors");
}

export function AdminGetBeds(){
  return api.get<Res<Bed[]>>("/admin/beds");
}

export function AdminAddDoctor(
  name:string,
  department_id:number,
) {
  return api.post<Res<null>>("/admin/doctor",{
    name,
    department_id,
  });
}

export function AdminDeleteDoctor(id:number){
  return api.delete<Res<null>>("/admin/doctor",{
    params:{id}
  });
}

export function AdminUpdateDoctor(
  id:number,
  department_id:number,
){
  return api.put<Res<null>>("/admin/doctor",{
    id,
    department_id,
  });
}

export function AdminAddBed(
  number: string,
  department_id: number,
){
  return api.post<Res<null>>("/admin/bed",{
    number,
    department_id,
  });
}

export function AdminDeleteBed(id:number){
  return api.delete<Res<null>>("/admin/bed",{
    params:{id}
  });
}

export function AdminGetAllPatients(){
  return api.get<Res<Patient[]>>("/admin/patients");
}

export function AdminAddPatient(
  request: {
    name:string,
    doctor_id:number,
    bed_id:number,
    in_time:string,
    state:string,
    gender:string,
    phone:string,
    birth:string,
  }
) {
  return api.post<Res<null>>(
    "/admin/patient",
    request
  );
}

export function AdminDeletePatient(id:number){
  return api.delete<Res<null>>("/admin/patient",{
    params:{id}
  });
}

export function AdminGetPatient(id:number){
  return api.get<Res<PatientDetail>>("/admin/patient/"+id);
}

export function AdminGetVitalSigns(
  bed_id:number,
  time_start:string | null,
  time_utill:string | null,
) {
  // console.log(time_start,time_utill);
  return api.get<Res<VitalSign[]>>(
    "/admin/vital_signs/"+bed_id, {
      params: {
        time_start,
        time_utill
      }
    }
  );
}

export function AdminUpdatePatient(
  id:number,
  name:string,
  in_time:string,
  state:string,
  gender:string,
  phone:string,
) {
  return api.put<Res<null>>("/admin/patient",{
    id,
    name,
    in_time,
    state,
    gender,
    phone,
  });
}
