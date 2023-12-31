import { Res } from "@type/response";
import { Role } from "@/type/user";
import {api} from "./api";

export function Login(id: number, password: string, role: Role) {
  return api.post<Res<null>>(`/login/${role}`, {
    id: id as number,
    password: password,
  });
}
