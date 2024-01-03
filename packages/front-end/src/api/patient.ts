import { api } from "./api";
import type { Patient, Res, VitalSign} from "@/type";

export function PatientGetInfo(): Promise<Res<Patient>> {
  return api.get("/patient/info");
}

export function PatientGetVitalSigns(
  time_start: string,
  time_until: string
): Promise<Res<VitalSign[]>> {
  return api.get("/patient/vital_signs", {
    params: {
      time_start,
      time_until,
    },
  });
}

export function PatientCall() {
  return api.get("/patient/call");
}
