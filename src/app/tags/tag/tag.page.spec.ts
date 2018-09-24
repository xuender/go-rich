import { CUSTOM_ELEMENTS_SCHEMA } from '@angular/core';
import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TagPage } from './tag.page';

describe('TagPage', () => {
  let component: TagPage;
  let fixture: ComponentFixture<TagPage>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TagPage ],
      schemas: [CUSTOM_ELEMENTS_SCHEMA],
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TagPage);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
