import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import { Observable} from 'rxjs';

export interface Service {
    id: number;
    name: string;
    duration: number; 
    price: number;
}

export interface Staff {
    id: number;
    name: string; 
}

@Injectable({ providedIn: 'root' })
export class ApiService {
  private baseUrl = 'http://localhost:8080';

  constructor(private http: HttpClient) {}

  getServices(): Observable<Service[]> {
    return this.http.get<Service[]>(`${this.baseUrl}/api/services`);
  }

  getStaff(): Observable<Staff[]> {
    return this.http.get<Staff[]>(`${this.baseUrl}/api/staff`);
  }
}