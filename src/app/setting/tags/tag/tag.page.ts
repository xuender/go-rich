import { Component } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';
import { HttpClient } from '@angular/common/http';

import { Tag } from '../../../api/tag';
import { URL } from '../../../api/init'
import { NgForm } from '@angular/forms';

@Component({
  selector: 'app-tag',
  templateUrl: './tag.page.html',
  styleUrls: ['./tag.page.scss'],
})
export class TagPage {
  tag: Tag
  constructor(
    private http: HttpClient,
    private modalCtrl: ModalController,
    private navParams: NavParams
  ) {
    this.tag = this.navParams.get('tag')
    if (!this.tag.type) {
      this.tag.type = {}
    }
  }

  cancel() {
    this.modalCtrl.dismiss();
  }

  save() {
    if (this.tag.id) {
      this.http.put<Tag>(`${URL}/api/tags/${this.tag.id}`, this.tag)
        .subscribe(tag => this.modalCtrl.dismiss(tag))
    } else {
      this.http.post<Tag>(`${URL}/api/tags`, this.tag)
        .subscribe(tag => this.modalCtrl.dismiss(tag))
    }
  }
}
