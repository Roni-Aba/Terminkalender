import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ApiService, Service, Staff } from './api.service';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule],
  template: `
    <div style="padding: 24px; font-family: system-ui">
      <h1>Salon Kalender</h1>

      <h2>Services</h2>
      <ul>
        <li *ngFor="let s of services">
          {{ s.name }} – {{ s.duration }} min –
          {{ s.price | number:'1.2-2' }} €
        </li>
      </ul>

      <h2>Mitarbeiter</h2>
      <ul>
        <li *ngFor="let m of staff">{{ m.name }}</li>
      </ul>
    </div>
  `,
})
export class App implements OnInit {
  services: Service[] = [];
  staff: Staff[] = [];

  constructor(private api: ApiService) {}

  ngOnInit(): void {
    this.api.getServices().subscribe((v) => (this.services = v));
    this.api.getStaff().subscribe((v) => (this.staff = v));
  }
}