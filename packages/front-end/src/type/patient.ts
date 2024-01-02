export interface Patient {
  id: number;
  name: string;
  doctor: {
    id: number;
    name: string;
  }
  bed: {
    id: number;
    number: string;
  }
  state: string;
  in_time: string;
  out_time: string;
}
// ID            uint64       `gorm:"primary_key" json:"id"`
// BedID         uint64       `json:"bed_id"`
// Time          utility.Time `json:"time"`
// Temperature   float64      `json:"temperature"`
// HeartRate     uint64       `json:"heart_rate"`
// BloodPressure string       `json:"blood_pressure"`
// BreathingRate uint64       `json:"breathing_rate"`
// BloodOxygen   float64      `json:"blood_oxygen"`
// Sense         string       `json:"sense"`
export interface VitalSign {
  id: number;
  bed_id: number;
  time: string;
  temperature: number;
  heart_rate: number;
  blood_pressure: string;
  breathing_rate: number;
  blood_oxygen: number;
  sense: string;
}

export interface PatientDetail extends Patient {
  gender: string;
  phone: string;
  birth: string;
}
