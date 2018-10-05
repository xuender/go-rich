import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TradesPage } from './trades.page';

describe('TradesPage', () => {
  let component: TradesPage;
  let fixture: ComponentFixture<TradesPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TradesPage ],
      schemas: [CUSTOM_ELEMENTS_SCHEMA],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TradesPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
