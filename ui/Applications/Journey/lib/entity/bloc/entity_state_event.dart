import 'package:flutter/material.dart';
import 'package:equatable/equatable.dart';

@immutable
abstract class EntityStateEvent extends Equatable {
  const EntityStateEvent();
}

class EntityStateStarted extends EntityStateEvent {
  @override
  List<Object> get props => [];
}