import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
@Component({
  selector: 'app-tabs',
  templateUrl: 'tabs.page.html',
  styleUrls: ['tabs.page.scss']
})
export class TabsPage implements OnInit {
  constructor(
    private router: Router,
  ) {
  }
  ngOnInit() {
    const token = localStorage.getItem('token')
    if (!token || token.length < 10) {
      this.start()
    }
  }
  private start() {
    this.router.navigateByUrl('/start')
  }
}
