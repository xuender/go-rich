import { Component, OnInit } from '@angular/core';
import { ModalController, NavParams } from '@ionic/angular';
import { Ext } from '../../api/ext';
import {remove} from 'lodash'

@Component({
  selector: 'app-ext',
  templateUrl: './ext.page.html',
  styleUrls: ['./ext.page.scss'],
})
export class ExtPage implements OnInit {

  exts: Ext[] = []
  constructor(
    private modalCtrl: ModalController,
    private navParams: NavParams
  ) {
    this.exts = this.navParams.get('exts')
  }

  ngOnInit() {
  }
  save() {
    this.modalCtrl.dismiss(this.exts)
  }
  cancel() {
    this.modalCtrl.dismiss();
  }
  remove(e: Ext){
    remove(this.exts, e)
  }
}
